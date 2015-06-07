def detect_anagrams(word, sentence):
	wordset = []
	wordarray = word.lower().replace(""," ").split()
	baseletters = {}
	for char in wordarray:
		if char in baseletters:
			baseletters[char] += 1
		else:
			baseletters[char] = 1
	# print word, baseletters
	# print "\n"

	for nextword in sentence:
		print nextword, word
		if nextword is not word:
			completters = {}
			for char in nextword.replace(""," ").split():
				if char in completters:
					completters[char] += 1
				else:
					completters[char] = 1
			# print completters, completters == baseletters
			if completters == baseletters:
				wordset.append(nextword)
	return wordset