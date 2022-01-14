#ifndef KHAIIIC_CORE_H
#define KHAIIIC_CORE_H

#include "khaiii_api.h"
#include "khaiii_dev.h"

#ifdef __cplusplus
extern "C" {
#endif

int* create(const char* rsc_dir, const char* opt_str);
const khaiii_word_t* generate_analyze(int* tagger, const char* line);
char* analyze_morphs(int* tagger, const khaiii_word_t* words);
char* parse(int* tagger, const char* line);
int destroy(int* tagger);

#ifdef __cplusplus
}
#endif

#endif    // KHAIIIC_CORE_H
