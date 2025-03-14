package main

var Facts = [100]string{
	"The first computer programmer was Ada Lovelace, who wrote the first algorithm for Charles Babbage's Analytical Engine in the 1840s.",
	"The term 'bug' to describe a software defect originated when Grace Hopper found a moth causing issues in the Harvard Mark II computer in 1947.",
	"Python was named after the British comedy group Monty Python, not the snake.",
	"The first computer virus was created in 1971 and was called the 'Creeper'. It displayed the message 'I'm the creeper, catch me if you can!'",
	"The QWERTY keyboard layout was designed to slow down typing to prevent mechanical typewriters from jamming.",
	"JavaScript was created in just 10 days by Brendan Eich in 1995.",
	"The first version of Apple's iOS didn't have copy and paste functionality.",
	"The programming language C was developed at Bell Labs between 1969 and 1973 by Dennis Ritchie.",
	"Linus Torvalds created Linux as a student project when he was just 21 years old.",
	"The first computer mouse was made of wood and was invented by Doug Engelbart in 1964.",
	"The world's first website, created by Tim Berners-Lee, went live on August 6, 1991.",
	"Ruby was created by Yukihiro Matsumoto, who wanted a programming language designed for programmer happiness.",
	"The 404 error code for webpage not found originated from a room number at CERN where the World Wide Web was developed.",
	"The first email was sent by Ray Tomlinson in 1971. It was a test message sent to himself.",
	"The MD5 hashing algorithm, once widely used for security, was effectively broken in 2005.",
	"SQL (Structured Query Language) was developed at IBM in the 1970s.",
	"The concept of object-oriented programming was introduced in the programming language Simula 67 in 1967.",
	"The first computer game, 'Spacewar!', was developed by MIT students in 1962.",
	"The most expensive bug in history was the Ariane 5 rocket explosion in 1996, caused by a software error that cost approximately $370 million.",
	"Google's name originated from the misspelling of 'googol', which is the number 1 followed by 100 zeros.",
	"The Git version control system was created by Linus Torvalds in 2005 for Linux kernel development.",
	"Java was originally called 'Oak' after an oak tree outside the developer James Gosling's office.",
	"The 'Hello, World!' program tradition began with Brian Kernighan's 1978 book on C programming.",
	"Unicode, which allows computers to handle text in virtually all of the world's writing systems, contains over 143,000 characters.",
	"The Y2K bug occurred because early programmers used two digits to represent the year to save memory space.",
	"The HTTP status code 418 'I'm a teapot' was an April Fools' joke from 1998.",
	"The first hard drive (IBM 350) from 1956 weighed over a ton and stored only 3.75 MB of data.",
	"Go (Golang) was developed at Google in 2007 by Robert Griesemer, Rob Pike, and Ken Thompson.",
	"The first graphical web browser, 'Mosaic', was released in 1993.",
	"Perl was originally developed for text manipulation, and its name stands for 'Practical Extraction and Reporting Language'.",
	"The Turing machine, a theoretical computing model, was proposed by Alan Turing in 1936, long before physical computers existed.",
	"Bitcoin's creator, known by the pseudonym Satoshi Nakamoto, has never been conclusively identified.",
	"Rust was named after a fungus that is tough, distributed, and parallel, just like the programming language aims to be.",
	"The first webcam was set up at Cambridge University to monitor a coffee pot, saving researchers from making unnecessary trips to an empty pot.",
	"The floating-point error in Intel's Pentium processor (1994) cost the company $475 million in recall costs.",
	"The concept of 'recursion' in programming occurs when a function calls itself to solve smaller instances of the same problem.",
	"The first commercial computer, UNIVAC I, was delivered to the US Census Bureau in 1951.",
	"The `:-)` emoticon was first proposed by Scott Fahlman in 1982 at Carnegie Mellon University.",
	"The Apollo 11 guidance computer (AGC) had less processing power than a modern calculator.",
	"There are over 700 different programming languages in existence, though only a few dozen are widely used.",
	"The C++ programming language was originally called 'C with Classes'.",
	"Assembly language is considered a low-level programming language because it works closely with the computer's hardware.",
	"The first mobile phone call was made by Martin Cooper of Motorola in 1973.",
	"Fortran, developed in the 1950s, is still used today, particularly in scientific and engineering applications.",
	"The term 'Wi-Fi' doesn't actually stand for anything. It was created by a marketing firm as a play on 'Hi-Fi'.",
	"Donald Knuth's 'The Art of Computer Programming' is one of the most respected references in computer science, though few have read it entirely.",
	"The longest-lived computer bug was in the Zune music player, causing devices to freeze on December 31, 2008, due to an error in leap year calculation.",
	"Approximately 70% of all COBOL transactions done in the USA happen on IBM mainframes.",
	"The '@' symbol was used in email addressing by Ray Tomlinson because it was 'the only preposition' on the keyboard.",
	"The average smartphone today has more computing power than all of NASA had during the Apollo moon mission.",
	"The SQL language has no 'official' pronunciation - people say both 'sequel' and 'S-Q-L'.",
	"The ZIP file format was created by Phil Katz in 1989.",
	"The first computer to defeat a world chess champion was IBM's Deep Blue, which beat Garry Kasparov in 1997.",
	"Internet Protocol version 6 (IPv6) allows approximately 340 undecillion (3.4×10^38) unique IP addresses.",
	"The first open-source software license, the GNU General Public License (GPL), was written by Richard Stallman in 1989.",
	"Moore's Law, stating that transistor counts double approximately every two years, was formulated by Gordon Moore in 1965.",
	"The concept of 'pair programming' originated in the early days of computing when computers were expensive and developers had to share machines.",
	"Approximately 70% of web servers run on some form of Unix/Linux operating system.",
	"The BASIC programming language was developed in 1964 to enable students in non-scientific fields to use computers.",
	"The Apple Lisa, released in 1983, was named after Steve Jobs' daughter.",
	"The Brainfuck programming language uses only eight commands and was designed to challenge programmers with its minimalism.",
	"The first video game console, the Magnavox Odyssey, was released in 1972.",
	"The term 'algorithm' comes from the name of Persian mathematician Muhammad ibn Musa al-Khwarizmi.",
	"The web server Apache got its name because it was 'a patchy' server, made from patches to the NCSA HTTPd server.",
	"About 12.9 million developers use JavaScript, making it the most popular programming language in the world.",
	"The original name for Java's mascot, Duke, was 'The Waving Man'.",
	"HTTP/3, the third major version of HTTP, uses QUIC (a UDP-based protocol) instead of TCP to reduce latency.",
	"The Commodore 64, released in 1982, is still the single best-selling personal computer model of all time.",
	"The average game of chess has more possible variations than there are atoms in the observable universe.",
	"The term 'Software Engineering' was first used at a NATO conference in 1968.",
	"The first version of Microsoft Windows (Windows 1.0) was released on November 20, 1985.",
	"The SHA-256 hashing algorithm, used in Bitcoin mining, requires more than 4 billion combinations to find a matching hash.",
	"The 'Agile Manifesto' for software development was formulated in 2001 by 17 developers at a ski resort in Utah.",
	"The world's first digital camera was built by Kodak engineer Steve Sasson in 1975. It weighed 8 pounds and captured 0.01 MP images.",
	"The 'Konami Code' (↑↑↓↓←→←→BA) was first implemented in the 1986 game Gradius due to the game being too hard for testers.",
	"The first computer antivirus program was created in 1971 and was called 'Reaper', designed to remove the Creeper virus.",
	"The original Macintosh computer had 128KB of RAM, less than what it takes to display one modern JPEG image.",
	"The Android operating system's logo (the green robot) is named 'Bugdroid'.",
	"The IBM PC model 5150, introduced in 1981, established many of the standards for personal computers that are still in use today.",
	"The world's first website is still online at info.cern.ch.",
	"The idea for the computer spreadsheet came to Dan Bricklin when he was watching his professor manually recalculate a budget table in 1978.",
	"The PDF file format was created by Adobe in 1993 and became an open standard in 2008.",
	"The transistor, a key component in modern electronics, was invented at Bell Labs in 1947.",
	"The first documented denial-of-service attack occurred in 1974 at the University of Illinois.",
	"Unicode's most recently added emoji required over 5000 lines of documentation for its implementation.",
	"LISP, one of the oldest programming languages still in use, was specified in 1958.",
	"The concept of 'hyperlinks' was first proposed by Vannevar Bush in 1945, long before computers could implement them.",
	"The average developer creates 70 bugs per 1000 lines of code and fixes 61 of them.",
	"The F# programming language is named so because F is the musical notation for the key that the C# language is 'transposed' into.",
	"The first Bitcoin transaction was for two pizzas, purchased for 10,000 BTC (worth billions today).",
	"The default Windows XP wallpaper, 'Bliss', is a real photograph of a hill in Sonoma County, California.",
	"Over 90% of the world's currency exists only on computers, not as physical cash.",
	"The Linux kernel, which powers everything from Android phones to supercomputers, consists of over 27.8 million lines of code.",
	"Approximately 94% of ATMs worldwide run on a version of Windows.",
	"The popular image format PNG was partly created in response to licensing issues with the GIF format.",
	"The 'thunk', a delayed computation technique, was invented for the ALGOL 60 programming language in the early 1960s.",
	"Vim, a text editor first released in 1991, was created as an improved clone of vi, which itself was written in 1976.",
}
