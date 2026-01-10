# file-indexer

Search your files by keyword. We'll load the files into a search index and then expose them via HTTP. Users will see a search bar that forwards their queries to our HTTP server.

Disclaimer - this project is for learning Go and probably won't be a showcase of best practices

Simple roadmap:
1. Create server
    - expose endpoint: GET search/query/{query}
	- server lives in a daemon thread you either:
		1. launch manually
		2. launch automatically on startup (probably less portable, more config)
2. Implement Index
    - take file names as input (cmd-line arg)
    - iterate through the files indexing them line-by-line
	- storage options:
		1. In memory (e.g. Map)
		2. Bleve index
		3. SQLite
3. Machine learning
    - would be cool to search files by fuzzy concepts like "tone"
    - this would pair well with story scripts and other creative stuff
	- HTTP: GET search/mood/{mood}
        - character names (and other narrowing filters) as query parameters
