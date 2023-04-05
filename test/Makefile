### vim adds a newline at end of file
### truncate -s -1 sample.txt

MD5_SIMPLE_JSON := $(shell sha256sum simple.json | cut -d' ' -f1)
MD5_SIMPLE_ENC := $(shell sha256sum simple.urlenc | cut -d' ' -f1)
MD5_SIMPLE_ENC_OURS := $(shell cat simple.json | ../urlencode | sha256sum | cut -d' ' -f1)
MD5_SIMPLE_BOTH := $(shell cat simple.json | ../urlencode | ../urldecode | sha256sum | cut -d' ' -f1)

MD5_MULTILINE_JSON := $(shell sha256sum multiline.json | cut -d' ' -f1)
MD5_MULTILINE_ENC := $(shell sha256sum multiline.urlenc | cut -d' ' -f1)
MD5_MULTILINE_ENC_OURS := $(shell cat multiline.json | ../urlencode | sha256sum | cut -d' ' -f1)
MD5_MULTILINE_BOTH := $(shell cat multiline.json | ../urlencode | ../urldecode | sha256sum | cut -d' ' -f1)


.PHONY: test
test:
	@echo "---simple test---"
	@echo "MD5_SIMPLE_JSON        $(MD5_SIMPLE_JSON)"
	@echo "MD5_SIMPLE_ENC         $(MD5_SIMPLE_ENC)"
	@echo "MD5_SIMPLE_ENC_OURS    $(MD5_SIMPLE_ENC_OURS)"
	@echo "MD5_SIMPLE_BOTH        $(MD5_SIMPLE_BOTH)"
	@echo "[0] testing simple encoding, passed: $(shell [ "$(MD5_SIMPLE_ENC)" == "$(MD5_SIMPLE_ENC_OURS)" ] && echo "TRUE" || echo "FALSE" )"
	@echo "[1] testing simple decoing,  passed: $(shell [ "$(MD5_SIMPLE_JSON)" == "$(MD5_SIMPLE_BOTH)" ] && echo "TRUE" || echo "FALSE" )"
	@echo "---multiline test---"
	@echo "MD5_MULTILINE_JSON     $(MD5_MULTILINE_JSON)"
	@echo "MD5_MULTILINE_ENC      $(MD5_MULTILINE_ENC)"
	@echo "MD5_MULTILINE_ENC_OURS $(MD5_MULTILINE_ENC_OURS)"
	@echo "MD5_MULTILINE_BOTH     $(MD5_MULTILINE_BOTH)"
	@echo "[2] testing MULTILINE encoding, passed: $(shell [ "$(MD5_MULTILINE_ENC)" == "$(MD5_MULTILINE_ENC_OURS)" ] && echo "TRUE" || echo "FALSE" )"
	@echo "[3] testing MULTILINE decoing,  passed: $(shell [ "$(MD5_MULTILINE_JSON)" == "$(MD5_MULTILINE_BOTH)" ] && echo "TRUE" || echo "FALSE" )"

setup_test:
	echo nope, do manually if needed..
	#fb588785782ece513139f7f443c150e30d4600cb552b36a7b180caa9c96a4c00  multiline.json
	#6ecb269018d185a09157c110fd764f6f67509b12f4713a376c256eec008ec0b3  multiline.urlenc
	#1930d385dd8c2479b9be7d9d5602478d99a7b7bddd556a6c411d87c0b24b6650  simple.json
	#158545ade6b4a6a53b1850dd0ef494d4526b288dd5eda776b0f81627c0d0a9b5  simple.urlenc
	#
	#cat test/multiline.json | ./urlencode > test/multiline.urlenc
	#cat test/simple.json | ./urlencode > test/simple.urlenc
