
#suit for all test cases
test:
	go test -coverprofile ./logic/cover.out -covermode=count ./logic

#clean output
clean:
	rm -rf *.jar
	rm -rf *.aar
	rm -rf ./logic/cover.out

setup:
	go get golang.org/x/mobile/cmd/gomobile
	gomobile init -ndk /home/ajay/Android/Sdk/ndk-bundle/

ANDROID_HOME := /home/ajay/Android/Sdk
export ANDROID_HOME

#build android sdk
android:
	$(call clean)
	export ANDROID_HOME
	gomobile bind -target=android -o logic.aar ./github.com/codewin_go/logic
