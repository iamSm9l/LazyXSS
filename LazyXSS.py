import os
import sys

with open("domains.txt", "r") as f:
    lines = f.readlines()

for line in lines:
    # Get all urls

    query = "gau " + line.strip("\n") + " >> dump.txt"
    os.system(query)
    # Get only the urls with = 
    os.system("./equalFilter/equalFilter >> haveParam.txt")
    #strip the vals
    os.system("./strip/strip >> stripped.txt")
    #get unique things only
    os.system("cat stripped.txt | sort | uniq >> unique.txt")
    # check if they exist in the db
    os.system("./checkdb/checkdb >> new.txt")
    # get original url with original values from haveParam.txt
    os.system("./getoriginal/getoriginal >> uniqueWithVals.txt")
    # Get only the ones that are online
    os.system("httpx -l uniqueWithVals.txt -fc 404,403,400,500,401,402 >> online.txt")
    # Fuzz each param
    #os.system("./LazyXSSfuzz")
    with open("domains.txt", "r+") as f:
        lineArray = f.readlines()
        f.seek(0)
        for i in lineArray:
            if i != line:
                f.write(i)
            f.truncate()

    with open("finished.txt", "a") as f:
        f.write(line)


    os.remove("dump.txt")
    os.remove("sorted.txt")
    os.remove("online.txt")