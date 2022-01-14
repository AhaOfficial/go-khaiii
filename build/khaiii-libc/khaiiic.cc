#include <khaiiic.h>
#include <khaiiic_core.h>
#include <stdio.h>

khaiii_model_t* Create(char* rsc_dir, char* opt_str)
{
    khaiii_model_t* model = (khaiii_model_t*)malloc(sizeof(khaiii_model_t));
    if ( model == NULL ) {
        fprintf(stderr, "[Error] Creating Model Error!");
        return NULL;
    }
    model->tagger = create(rsc_dir, opt_str);
    return model;
}

char* Parse(int* tagger, char* line)
{
    char* parsedResult = parse(tagger, line);
    if ( parsedResult == NULL ) {
        fprintf(stderr, "[Error] Parsing Sentence Error!");
    }
    return parsedResult;
}

int Destroy(khaiii_model_t* model)
{
    int is_detroyed = destroy(model->tagger);
    if ( !is_detroyed ) {
        fprintf(stderr, "[Error] Destroy Model Error!");
    } else {
        model->tagger = NULL;
        free(model);
        model = NULL;
    }
    return is_detroyed;
}
