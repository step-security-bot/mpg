# mpg

mpg is a command-line tool that generates strings of random characters
that can be used as reasonably secure passwords.

Passwords are, by default, chosen from the union of three character
classes: upper-case letters, lower-case letters, and digits.

Options can be given to omit any one or any two of these character
classes. For instance, you can omit uppercase letters and digits by
passing `-upper=false -digit=false` to `mpg`. This will return a
password composed only of lowercase letters.

Passwords are guaranteed to contain at least one of each selected
character class. The default length is 16. mpg will create a password
of any length greater than or equal to the number of character classes
selected.

If the password length is less than or equal to the total number of
characters selected, mpg will not repeat characters within the
generated password.
