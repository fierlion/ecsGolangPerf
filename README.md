# ecsGolangPerf
basic server and docker stuff to test ecs performance

```
docker build -t perf-server .
docker images --filter reference=perf-server
docker run -i -t -d -p 80:80 server
```

`http://ec2-<some ec2>.<region>.compute.amazonaws.com/?key=<the text to hash>&count=<the number of times (int) you'd like the hash to run>`
ie
`http://ec2-00-00-00-00.us-west-2.compute.amazonaws.com/?key=test&count=3`

login to ecr
```
$eval $(aws ecr get-login --no-include-email --region us-west-2 | sed 's|https://||')
```
then create a repo and push the image to ecr

```
aws ecr create-repository --repository-name perf-server-repository --region <region>
```
