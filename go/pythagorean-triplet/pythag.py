from math import sqrt

mn = 1
mx = 10
store = set()

for a in range(mn,mx):
	for b in range(mn,mx):
		c = sqrt((a**2)+(b**2))
		if c % 1 == 0:
			if  mn <= c <= mx:
				store.add((a, b, int(c)))

print store
# TODO: reduce set to unique sets