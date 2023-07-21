VERSION:=$(shell go version)
$(info $(VERSION))
ifneq (,$(findstring windows, $(VERSION)))
EXT:=.exe
else
ifneq (,$(findstring linux, $(VERSION)))
else
ifneq (,$(findstring darwin, $(VERSION)))
EXT:=.app
else
endif
endif
endif

huarongdao_init$(EXT): game/*.go init/*.go
	cd init && go build -o ../huarongdao_init$(EXT) && cd ..

statics/favicon.ico: images/favicon.png
	magick images/favicon.png -resize 256x256 - | magick - statics/favicon.ico

data/data: huarongdao_init$(EXT)
	./huarongdao_init$(EXT)

doserial$(EXT): serial/*.go
	cd serial && go build -o ../doserial$(EXT) && cd ..

generated/*.go: doserial$(EXT) statics/game.js statics/libcanvas.js statics/index.html statics/favicon.ico data/data
# go:embed is good, but it does not work on low version(1.10) golang, which is required for windowsxp
#bash -c 'perl -e '"'"'for $$a (@ARGV){open $$fh,"<", $$a or die $$!; @c=<$$fh>; $$data{$$a}=join "", map {if(ord($$_)>=32 && ord($$_)<=126 && ord($$_)!=34 && ord($$_)!=92 ){ $$_; }else{sprintf "\\x%02X", ord; } } split//, join "", @c; close $$fh; $$nn=$$a; $$nn=~s/\//_slash_/; $$nn=~s/\./_dot_/; $$namemap{$$a}=$$nn; open $$fh, ">", "generated/".$$nn."_content.go" or die $$!; print $$fh "package generated\nconst ", $$nn, " = \"", $$data{$$a}, "\"\n"; close $$fh;}; $$n="generated/contents.go"; open $$fh, ">", $$n or die $$1; print $$fh "package generated\nvar Contents = map[string]string{\n";       for $$k (keys %namemap){print $$fh "\"", $$k, "\":", $$namemap{$$k}, ",\n"};        print $$fh "}\n"; close $$fh; '"'"" statics/game.js statics/libcanvas.js statics/index.html statics/favicon.ico data/data "
	./doserial$(EXT) statics/game.js statics/libcanvas.js statics/index.html statics/favicon.ico +data/data

huarongdao-darwin/huarongdao-oc/libhuarongdao-arm64.a: go.mod go.sum game/*.go generated/*.go
	CGO_ENABLED=1 GOOS=darwin GOARCH=arm64 go build -o huarongdao-darwin/huarongdao-oc//libhuarongdao-arm64.a -buildmode c-archive && \
	cp huarongdao-darwin/huarongdao-oc/libhuarongdao-arm64.h huarongdao-darwin/huarongdao-oc/libhuarongdao.h

huarongdao-darwin/huarongdao-oc/libhuarongdao-amd64.a: go.mod go.sum game/*.go generated/*.go
	CGO_ENABLED=1 GOOS=darwin GOARCH=amd64 go build -o huarongdao-darwin/huarongdao-oc/libhuarongdao-amd64.a -buildmode c-archive && \
	cp huarongdao-darwin/huarongdao-oc/libhuarongdao-amd64.h huarongdao-darwin/huarongdao-oc/libhuarongdao.h

huarongdao-darwin/huarongdao-oc/libhuarongdao.a: huarongdao-darwin/huarongdao-oc/libhuarongdao-arm64.a huarongdao-darwin/huarongdao-oc/libhuarongdao-amd64.a
	lipo -create huarongdao-darwin/huarongdao-oc/libhuarongdao-arm64.a huarongdao-darwin/huarongdao-oc/libhuarongdao-amd64.a -o huarongdao-darwin/huarongdao-oc/libhuarongdao.a

statics/*.js: typescripts/src/*
	cd typescripts && tsc && cd ..

huarongdao: go.mod go.sum game/*.go generated/*.go *.go
	go build

rsrc_windows_386.syso: huarongdao.exe.manifest statics\favicon.ico
	rsrc -manifest huarongdao.exe.manifest -ico ./statics/favicon.ico -arch 386

huarongdao.exe: go.mod go.sum game/*.go generated/*.go *.go rsrc_windows_amd64.syso
	go build -ldflags "-H=windowsgui"

huarongdaoxp.exe: go.mod go.sum game/*.go generated/*.go *.go rsrc_windows_386.syso
	rename start_notxp.go start_notxp.go.tmp
	rename browser_search_notxp_windows.go browser_search_notxp_windows.go.tmp
	go build -ldflags "-H=windowsgui" -tags windowsxp -o huarongdaoxp.exe
	rename browser_search_notxp_windows.go.tmp browser_search_notxp_windows.go
	rename start_notxp.go.tmp start_notxp.go

huarongdao.app.zip: $(info $(wildcard huarongdao-darwin/**/*)) go.mod go.sum game/*.go generated/*.go *.go
	cd huarongdao-darwin && xcodebuild -configuration Release && cd Build/Release && zip -r ../../../huarongdao.app.zip huarongdao-oc.app && cd ..
