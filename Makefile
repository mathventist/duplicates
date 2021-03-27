all: pp str2sentences str2ngrams compareEqual compareWord2Vec containment resemblance

pp:
	go build ./cmd/pp

str2sentences:
	go build ./cmd/str2sentences

str2ngrams:
	go build ./cmd/str2ngrams

compareEqual:
	go build ./cmd/compareEqual

compareWord2Vec:
	go build ./cmd/compareWord2Vec

containment:
	go build ./cmd/containment

resemblance:
	go build ./cmd/resemblance

clean:
	rm -f pp str2sentences str2ngrams compareEqual compareWord2Vec containment resemblance
