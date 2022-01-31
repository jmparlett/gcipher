# Gocipher Notes
### Created: 2021-December-30

## Ceaser
First cipher to implement will just be ceaser. Want to implement encryption and decryption given a key
and brute forcing.

Assume *k* is out key value.

Ceaser cipher/rot is pretty simple. For encryption add *k%26*. For decryption subtract *k%26*.
We will act only on alphabetical chars. We could add special cases for rot13 and stuff, but its
a waste of time. With rot weather or not you add or subtract you get the original val back.

For semantics, we'll do all encryption and decryption in the world of strings. We're not going to have
the cipher libs care about things like file handling, arg parsing, etc. They get a string and based on the
call they either encrypt, decrypt, or attempt to brute force.

## Rot algos
Rot is a little easier or a little odder. 
- rot13 is just the ceaser cipher with a key of 13
- rot5 rotates only numbers by 5
- rot 18 is the application of both rot13 and rot 5 to the text in question
- rot 47 rotates any char in the range [33(!)-126(~)] 47 places forward. So the size of the rotation space is expanded
, but the process is still basically the same.

`chr(33+((ord('2')+14)%94))` This should be useable for any rot cipher. 94 is k in this case (rot47) 47\*2=94

33 is start of range 93+33=126 is the end of the range. If we take mod 94 we guarantee it is at most 93 so we never run over the range
If we are 94 and above then we reset to the begining of the range. 

Remember rots are symmetrical in transformation if you take something 47 places ahead in the range then 47 places again you've 
gone in a circle.


```
s=start
e=end
c=char
l=len of range
k=shift

Rot formula
s + ((c+k-s)%l)

```

If I want to stay in a range between 0 and 10 when I am shifting by 5 I simply take mod table size (x+5)%10

(0+5)%10 = 5
(1+5)%10 = 6
(6+5)%10 = 1

You'll never leave the range

What if I want to stay in a range between 10 and 20, (x+5)%20 when x=20 takes us out of the range

If its always a range of 10 numbers you can just use the old calculation and add 10 to the result 10+((x+5)%10)

So what about a range of 40 numbers say 50-90, from the previous result its just 50+((x+20)%40)

50+((x+20-50)%40)

## Substituiton Cipher
The goal of this cipher will be take the entire ascii table to another permutation and vice versa. So we need a way generate a random permutation of
numbers from 0-255.

Substituiton ciphers are good against brute forcing because for an alphabet of size *n* there are *n!* possible permutations to check. They are
vulnerable to frequency analysis however because the frequency of letters will match that of the alphabet for a large enough cipher text.

What if you actually transformed it into an alien language first then took permutations of the alphabet to translate it? Would it still be vulnerable?

### Permutation
We need to generate a permuation and its inverse.

p1 is the inverse of p2 and vice versa.

*p1	=	{3,8,5,10,9,4,6,1,7,2}*

*p2	=	{8,10,1,6,3,7,9,2,5,4}*

This is because we can use all the positions in *p2* as indices into *p1* to get the orginal numbers in order from 1 to 10 and vice versa for
*p1* as indices for *p2*.

To generate *p2* (or any inverse permutation)
```python
for i in p1:
  p2[p1[i]]=i
```

We can use any the below algorithms to generate random permutation. Essentially your just performing n random swap operations on an array
to create a new mapping. 

![perm algos](randomperms_algo.png)

Taken from "Generating Random Permutations", a PHD thesis by Jorg Arndt.


Since it is a random number generator we can actually use the seed as the key to decrypt our cipher text.

### Encryption
1. Create an array of len *256*
2. Generate a random number (the key)
3. Seed the random number generator with this value
4. perform *256* random swaps in the array
5. generate the cipher text by using each char as an index to the array adding a char to the ciphertext each time
6. return cipher text
7. return random seed

### Decryption
1. Create an array of len *256*
2. Use the provided key to seed the random number generator
3. Perform *256* random swaps
4. Generate the inverse of the created permutation
5. generate the plaintext by using each char as an index to the array adding a char to the plaintext each time
6. return the plaintext


## Vigenere Cipher
The goal of this cipher is to take a keyword of length *n*. Each char will correspond to a shift in a sequence. Then *n* ceaser shifts will
be applied in sequence to generate the given cipher text.


## Affine Cipher
The goal of this cipher is to take a keyword that represents a shift and 
