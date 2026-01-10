# file-indexer

Simple program to search for keywords in a list of files. We'll load the files into a search index and support queries via HTTP. The UI will just be a simple search bar that queries our server.

Disclaimer - this project is for learning Go and probably won't be a showcase of best practices

Roadmap:
1. Create HTTP server
    - create server-side endpoints
    - implement handlers that grab results from the index
	- launch server on-demand
		1. manually via command line?
		2. automatically via service registry?
2. Create search index
    - take the list of file names as program input (eg. cmd line arg)
    - iterate through the files & index them line-by-line
	- storage options (based on persistence needs):
		1. In memory (eg. Map)
		2. Bleve index
		3. SQLite
3. Experiment with fuzzy search
    - would be interesting to search by fuzzy concepts like tone, synonyms, etc.

