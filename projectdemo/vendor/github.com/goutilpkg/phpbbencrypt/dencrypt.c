#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <math.h>

#include "md5.h" 

#define max(a,b)(a>b?a:b)

#define min(a,b)(a<b?a:b)

typedef unsigned char byte;

char *ITOA64 = "./0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz";

int _hash_crypt_private(char *hash_output, char *password, char *setting, char *itoa64);

int md5(char* str, char *hex_output, int len) {
    md5_state_t state;
    md5_byte_t digest[16] = {0};
    int di;

    md5_init(&state);
    md5_append(&state, (const md5_byte_t *)str, len);
    md5_finish(&state, digest);
    for (di = 0; di < 16; ++di)
        sprintf(hex_output + di * 2, "%02x", digest[di]);
    return 0;
}

int unique_id(char *id) {
    char timestr[16 * 2 + 1] = {0};
    sprintf(timestr, "%ld", time(0));
    char md5_output[16 * 2 + 1];
    md5(timestr, md5_output, strlen(timestr));
    md5_output[20]=0;
    strcpy(id, &(md5_output[3]));
    return 0;
}

int _hash_encode64(char *output, byte *text, int count, char *itoa64) {
    int i = 0;
    if (i > count) {
        i = count;
    }

    while (1) {
        if (i >= count) {
            break;
        }
        int value = (unsigned int)text[i];
        *output++ = itoa64[value & 0x3f];
        i += 1;

        if (i < count) {
            value |= (int)text[i] << 8;
        }
        *output++ = itoa64[(value >> 6) & 0x3f];

        if (i >= count) {
            break;
        }
        i += 1;

        if (i < count) {
            value |= (int)text[i] << 16;
        }

        *output++ = itoa64[(value >> 12) & 0x3f];

        if (i >= count) {
            break;
        }
        i += 1;

        *output++ = itoa64[(value >> 18) & 0x3f];
    }
    *output = 0;
    return 0 ;
}

int _hash_gensalt_private(char* output, char *text, char *itoa64, int iteration_count_log2) {
    if ((iteration_count_log2 < 4) || ( iteration_count_log2 > 31)){
        iteration_count_log2 = 8;
    }
    char slat[4+10+1] = {0};
    strcpy(slat, "$H$");
    slat[3] = itoa64[min(iteration_count_log2 + 5, 30)];
    //slat[3] = '6';
    _hash_encode64(&(slat[4]), (byte*)text, 6, itoa64);
    strncpy(output, slat, 14);
    return 0;
}

int phpbb_hash(char *output, char *password, char *itoa64) {
    char random_state[16+1];
    unique_id(random_state);
    char random[6+1] = {0};
    strncpy(random, random_state, 6);
    int count = 6;

    char slat[6 + 20  +1] = {0};
    _hash_gensalt_private(slat, random, itoa64, 6); 
    char hash_output[16*2+1] = {0};
    _hash_crypt_private(hash_output, password, slat, itoa64);

    if (strlen(hash_output) != 34) {
        md5(hash_output, password, strlen(password));
    }
    strcpy(output, hash_output);
    return 0;
}

int hex2dec(char c) {
    switch (c) {
        case '0':
        case '1':
        case '2':
        case '3':
        case '4':
        case '5':
        case '6':
        case '7':
        case '8':
        case '9':
            return c-'0';
        case 'a':
        case 'b':
        case 'c':
        case 'd':
        case 'e':
        case 'f':
            return c - 'a' + 10;
        default:
            return 0;    
    }
}

int unhexlify(char *output, char *hash) {
    while (*hash) {
        if (0 ==*(hash+1)) {
            return 1;
        }
        *output++ = (hex2dec(*hash)<<4) + hex2dec(*(hash+1));
        hash += 2;
    };
    *output = 0;
    return 0;
}

int _hash_crypt_private(char *hash_output, char *password, char *setting, char *itoa64) {
    *hash_output = '*';

    //    # Check for correct hash
    if (setting != strstr(setting, "$H$")) {
        return 0;
    }

    //count_log2 = itoa64.index(setting[3])
    int count_log2 = 0; 
    char *p = itoa64;
    for(; *p != 0; ++p) {
        if (setting[3] == *p) {
            count_log2 = (p - itoa64)/sizeof(*p);
        }
    }

    if ((count_log2 < 7) || (count_log2 > 30)){
        return 0;
    }

    int count = 1 << count_log2;
    // TODO: molloc memory here
    char salt[8+1024+1]= {0};
    strncpy(salt, &(setting[4]), 8);

    if (strlen(salt) != 8) {
        printf("salt len error");
        return 0; 
    }
    int plen = strlen(password);
    char hash[16*2 + 1] = {0};
    strncpy(&(salt[8]), password, 1024);
    md5(salt, hash, plen + 8);
    char hex[16+1+1024] = {0};
    unhexlify(hex, hash);

    strncpy(&(hex[16]), password, 1024);
    md5(hex, hash, plen + 16);
    unhexlify(hex, hash);
    count -= 1;
    while (count) {
        strncpy(&(hex[16]), password, 1024);
        md5(hex,hash, plen+16);
        unhexlify(hex, hash);
        count -= 1;
    }
    strncpy(hash_output, setting, 12);
    hash_output += 12;
    _hash_encode64(hash_output, (byte*)hex, 16, itoa64);
    return 0;
}

int phpbb_check_hash(char *password, char *hashtext) {
    if (34 == strlen(hashtext)) {
        char hash_output[34 + 1]; 
        _hash_crypt_private(hash_output, password, hashtext, ITOA64); 
        return 0 == strcmp(hash_output, hashtext);
    }
    char hex_output[16*2 + 1];
    md5(password, hex_output, 32);
    return 0 == strcmp(hex_output, hashtext);
}


int encrypt(char *hash_output, char* password) {
    return phpbb_hash(hash_output, password, ITOA64);
}

int verify(char *echopw, char * hash) {
    return phpbb_check_hash(echopw, hash);
}
