# Publisher

A cli tool that publishes an article to several platforms. Currently, supported platforms is medium.io and dev.to.

## File Format

The application expects an article in the following format. 

```markdown
    # Header
    
    ## Meta
    
    tags: go, programming

    ## Part 1, Intro
```

The level one header (`#`) will be used at a title for the article. Information in section `## Meta` represents meta information about your article and won't be included into its text. `Tags` represent a list of tags of the article.

Next section after (after ##Meta) will represent a body of the article.

## Usage

Usage: publisher [--devtoapikey DEVTOAPIKEY] [--mediumapikey MEDIUMAPIKEY] --platform PLATFORM --action ACTION [--filepath FILEPATH]

Options:

    --devtoapikey DEVTOAPIKEY   API key for dev.to platform [env: DEVTOAPIKEY]
    --mediumapikey MEDIUMAPIKEY API key for medium.io platform [env: MEDIUMAPIKEY]
    --platform PLATFORM    Platform to publish an article. Currently supported medium.io and dev.to
    --action ACTION        Action to perform: list or publish
    --filepath FILEPATH    Filepath to an article for publishing
    --help, -h             display this help and exit
