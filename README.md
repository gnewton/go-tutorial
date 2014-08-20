go-tutorial
===========

Code use in short (2hr) Go tutorial for some of MBB.

Read Fasta file and push each entry to a channel. 
Take entries from channels and give to 12 workers to perform computation on each.
 Right now only prints out entry.

Fasta reader not complete:
 * Does not handle `;` starting lines or those that wrap additional lines using `;`
 * Does not fast fail
See https://en.wikipedia.org/wiki/FASTA_format
