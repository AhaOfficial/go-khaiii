#include <khaiiic_core.h>
#include <KhaiiiApi.hpp>

std::shared_ptr<khaiii::KhaiiiApi> global_tagger;

int* create(const char* rsc_dir, const char* opt_str)
{
    try {
        global_tagger = khaiii::KhaiiiApi::create();
        global_tagger->open(rsc_dir, opt_str);
        return (int*)global_tagger.get();
    } catch (const khaiii::Except& exc) {
        return nullptr;
    }
}

const khaiii_word_t* generate_analyze(int* tagger, const char* line)
{
    const khaiii_word_t* results = nullptr;
    khaiii::KhaiiiApi* khaiii_api = (khaiii::KhaiiiApi*)tagger;
    try {
        results = khaiii_api->analyze(line, "");
        return results;
    } catch (const khaiii::Except& exc) {
        return nullptr;
    }
}

char* analyze_morphs(int* tagger, const khaiii_word_t* words)
{
    khaiii::KhaiiiApi* khaiii_api = (khaiii::KhaiiiApi*)tagger;
    char* result = nullptr;
    int result_size = 0;
    std::string strResult;
    const std::string sep = "/";
    const std::string token = " + ";

    for (auto word = words; word != nullptr; word = word->next) {
        const khaiii_morph_t* morphs = word->morphs;
        for (auto morph = morphs; morph != nullptr; morph = morph->next) {
            strResult += morph->lex;
            strResult += sep;
            strResult += morph->tag;

            if (morph->next != nullptr) {
                strResult += token;
            }
        }
        if (word->next != nullptr) {
            strResult += token;
        }
    }
    khaiii_api->free_results(words);

    result_size = sizeof(char) * strResult.size();
    if ( (result = (char*)malloc(result_size + 1)) == nullptr ) {
        return nullptr;
    } else {
        memset(result, 0x00, result_size + 1);
        memcpy(result, strResult.c_str(), result_size + 1);
    }
    return result;
}

char* parse(int* tagger, const char* line)
{
    char* result = nullptr;

    const khaiii_word_t* words = generate_analyze(tagger, line);
    if ( words == nullptr ) {
        return nullptr;
    }

    result = analyze_morphs(tagger, words);
    if ( result == nullptr ) {
        return nullptr;
    }

    return result;
}

int destroy(int* tagger)
{
    khaiii::KhaiiiApi* khaiii_api = (khaiii::KhaiiiApi*)tagger;
    try {
        khaiii_api->close();
        khaiii_api = nullptr;
        return 1;
    } catch (const khaiii::Except& exc) {
        return 0;
    }
}
