size_t go_write_callback(char *, size_t, size_t, void *);

static inline CURLcode
my_set_write(CURL *c_handle, void *go_handle) {
	CURLcode res;
	res = curl_easy_setopt(c_handle, CURLOPT_WRITEFUNCTION, go_write_callback);
	if (res != CURLE_OK)
		return res;
	res = curl_easy_setopt(c_handle, CURLOPT_WRITEDATA, go_handle);
	if (res != CURLE_OK)
		curl_easy_setopt(c_handle, CURLOPT_WRITEFUNCTION, NULL);
	return res;
}
