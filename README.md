# file-indexer

A Go program that indexes local files and then exposes the index for querying through a simple HTTP API. You just pass in a search term and get back a list of "hits" (file name, line number, line text). 

Simple roadmap:
1. Create HTTP server
    - expose endpoint: GET search/query/{query}
	- server lives in a daemon thread you either:
		1. launch manually
		2. launch automatically on startup (probably less portable, more config)
2. Index impl
    - take files or directories as input (cmd-line args?)
    - iterate through the file list indexing them line-by-line
	- storage options:
		1. brute-force (keyword->line Map)
		2. Bleve index
		3. SQLite
3. Machine learning
    - it would be cool to search files by fuzzy concepts like "tone"
    - this would pair well with story scripts, business emails, etc.
	- HTTP: GET search/mood/{mood}
        - "character_name" and other narrowing filters as query parameters


This is just for fun; I have no idea what I'm doing ;)