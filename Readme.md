# Ffuf: Post processing
Unfortunately - despite its "-ac" flag, ffuf tends to produce a lot of irrelevant entries. This is why I've created a post-processing tool to filter out those entries. 
Additionally, I saw a lot of relevant entries removed when "-ac" was used - especially when ffuf encounters json or xml files - "-ac" might drop them entirely.
This tool has to be run after ffuf has finished. Additionally, the initial ffuf command should be run with the following flags:

```
-o /folder/to/results.json
-od /folder/to/bodies
-of json (default)
```

This forces ffuf to write a summary file in json format as well as bodies of the responses to disk. 
Adding "-od" is not mandatory but recommended.

I highly recommend __NOT to use the "-ac" flag__ - especially if you don't want to miss cool stuff and want to use this
post-processing tool.

## Usage

```
Usage of ./ffufPostprocessing:
  -result-file string
        Path to the original ffuf result file (in json format)
  -bodies-folder string
        Path to the ffuf bodies folder (optional, if set results will be better)
  -new-result-file string
        Path to the new ffuf result file (optional)
  -delete-bodies
        Delete unnecessary body files after filtering (optional)
  -overwrite-result-file
        Overwrite original result file (optional)
  -verbose
        Verbose mode (Shows currently filtered results) (optional)
```

## Example

First run ffuf as always - I used a very simple minimalistic command:

```
gau example.net | unfurl -u paths
```
![image](https://github.com/reewardius/ffufPostprocessing/assets/68978608/9498f697-c19e-413a-b7ee-52b5131df10e)
```
sed 's#/#\n#g' paths.txt | sort -u
```
![image](https://github.com/reewardius/ffufPostprocessing/assets/68978608/f917316a-ff16-4335-b11f-5c74cc5322c8)
```
gau example.net | unfurl -u keys
```
![image](https://github.com/reewardius/ffufPostprocessing/assets/68978608/ae9db592-111e-41a5-a22e-36df7837d252)

![image](https://github.com/reewardius/ffufPostprocessing/assets/68978608/fab4d847-96a9-45a2-978c-d98af1f7a472)

![image](https://github.com/reewardius/ffufPostprocessing/assets/68978608/1ba1bed0-cfa7-4485-b74d-a6ab8d01615f)
```
curl https://uk.yahoo.com/ -o yahoo.html
cat yahoo.html | tok | tr '[:upper:]' '[:lower:]' | sort -u > yahoo-words
```
![image](https://github.com/reewardius/ffufPostprocessing/assets/68978608/744c8921-4b4f-4ddb-b0af-02372a824369)

Fetch JS
```
gau example.net | fff -s 200 -s 404 --output yahoo
find yahoo/ -type f -name '*.body' | html-tool attribs src | grep '\.js$'
```
![image](https://github.com/reewardius/ffufPostprocessing/assets/68978608/262fdbad-76a7-4d07-b6c2-388e97953f00)

Ffuf + FfufPostProcessing + GAP
![image](https://github.com/reewardius/ffufPostprocessing/assets/68978608/319cc98d-e7cb-482d-8832-76b4ee6afcbb)
![image](https://github.com/reewardius/ffufPostprocessing/assets/68978608/2709d373-bb42-46b9-bfc0-b19047707ae0)
![image](https://github.com/reewardius/ffufPostprocessing/assets/68978608/fddc9649-f422-4a9b-9229-b8c9e7d0f277)
![image](https://github.com/reewardius/ffufPostprocessing/assets/68978608/a6a82621-18d1-4cbf-bd06-55cd4525b8f3)
![image](https://github.com/reewardius/ffufPostprocessing/assets/68978608/36220145-6077-44bb-a322-841da2a5aa79)

Once GAP finishes. Copy out your links and feed them into file that you will run FFUF against. In my case, I filtered by "/api", "/admin" and "/user". Those are usually juicy.
```
./ffuf -w "GAPoutput.txt" -u "FUZZ" -noninteractive -o "/tmp/results.json" -od "/tmp/bodies/" -of json
./ffufPostprocessing -result-file "/tmp/results.json" -bodies-folder "/tmp/bodies/" -delete-bodies -overwrite-result-file
jq -r '.results[].url' "/tmp/results.json" | httpx -title -sc -lc -nc -silent | sed 's/[][]//g' | awk '$NF > 20' | egrep '200|301|302'
jq -r '.results[].url' "/tmp/results.json" | nuclei -tags token,tokens -es unknown
```


## Details

Especially when -od is set, which means we have all http headers and bodies for each requested url - this tool will initially
analyse all bodies and enrich the initial results json file with the following data points:

- count of all headers
- domain of redirect if applicable
- amount of parameters in redirect if applicable
- length and words of page title (if existent)
- count of detected css files
- count of detected js files
- count of tags in html/xml/json (calculation is wild)

Afterwards it will scan the entire new results file and keep only those entries which are unique based on known metadata types.
If it turns out that one of those values is always different (e.g. the title of pages can vary very much) - this metadata type 
will be skipped for the uniqueness check.

In general this tool will always keep a small amount of entries which are _not_ unique. For example, if the results json file
contains 300x http status 403 (with words, length, ... identical) and 2 unique http status 200 responses, it won't drop all 300 http status 403 entries. 
It will keep X of them in the data set.

## Install

ffufPostprocessing requires golang 1.19

### Build from source

```
cd ffufPostprocessing
go build -o dist/ffufPostprocessing main.go
```

## License

I don't care. Do whatever you want with this tool.
