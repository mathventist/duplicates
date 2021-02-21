all: pp str2sentences str2ngrams compareEqual containment resemblance

pp:
	go build ./cmd/pp

str2sentences:
	go build ./cmd/str2sentences

str2ngrams:
	go build ./cmd/str2ngrams

compareEqual:
	go build ./cmd/compareEqual

containment:
	go build ./cmd/containment

resemblance:
	go build ./cmd/resemblance

clean:
	rm -f pp str2sentences str2ngrams compareEqual containment resemblance
