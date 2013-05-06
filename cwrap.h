size_t go_write_callback(char *, size_t, size_t, void *);

static inline CURLcode
my_set_write(CURL *handle) {
	return curl_easy_setopt(handle, CURLOPT_WRITEFUNCTION, go_write_callback);
}
