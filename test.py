def fib(n):
	if n < 2:
		return n
	else:
		return fib(n-1) + fib(n-2)

print(fib(11))
#for i in range(100000):
#	print(i)
