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
