#ifndef KHAIIIC_H
#define KHAIIIC_H

#include <stdlib.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef struct khaiii_model_t_ {
    int *tagger;
} khaiii_model_t;

khaiii_model_t* Create(char* rsc_dir, char* opt_str);
char* Parse(int* tagger, char* line);
int Destroy(khaiii_model_t* model);

#ifdef __cplusplus
}
#endif

#endif    // KHAIIIC_H
