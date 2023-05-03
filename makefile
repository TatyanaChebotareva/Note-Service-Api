PHONY: generate
generate:
#	rm    pkg\note_v1\\

#	mkdir pkg\note_v1

#	protoc 	--proto_path api/note_v1	\
			--go_out=pkg\note_v1 --go_opt=paths=import \
			--go-grpc_out=pkg\note_v1 --go-grpc_opt=paths=import \
			api\note_v1\note.proto

# for some reason PS doesn't see capital letters. Also these two commands doesn't work from makefile :(
#	mv pkg\note_v1\github.com\tatyanachebotareva\note-service-api\pkg\note_v1\\* pkg\note_v1\\

#	rmdir pkg\note_v1\github.com