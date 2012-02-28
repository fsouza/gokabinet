#include <kclangc.h>
#include "kcgo.h"

void
free_pair(_pair p) {
	if (p.key != NULL) {
		free(p.key);
		p.key = NULL;
	}
}

_pair
gokccurget(KCCUR *cur) {
	_pair p;
	size_t ksiz, vsiz;
	const char *argvbuf;
	p.key = kccurget(cur, &ksiz, &argvbuf, &vsiz, 1);
	p.value = argvbuf;
	return p;
}
