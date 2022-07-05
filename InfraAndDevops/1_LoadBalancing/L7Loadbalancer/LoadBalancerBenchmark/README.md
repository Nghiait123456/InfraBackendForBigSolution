/**
Base code from link: https://github.com/gaplo917/load-balancer-benchmark
I have many problem, i custom for this form this base code

1) Load test local to local:
   1) context: same to author, you create 3 nodejs instance for backend load test, you create other tool L7LB : apache-event, apache-prework, apache-worker, haproxy, nginx, traefik in same docker-compose. All instance run in same docker-network.
   2) follow request: tool loadtest ==> L7LB  => nodejs.
   3) this is perfect context, every thing in local network, every is very fast and tiny. We will return {"200:OK"},  you need go to README_AUTHOR.md and run line 53 to 71.
   4) result overview: 
      1) in context local ==> local or local => endpoint google, we get result same 1_LoadBalancing/L7Loadbalancer/LoadBalancerBenchmark/result easy with cpu i5 8 core 8GB ram
      2) in context product, this is point we care important?
      3) you are very fast have quick test in local, but in product, it's not simple. <br/>
2) Load test product: We find answer for q/a: how many maximum rqs in one instance L7LB in production? 
   1) It's problem have many problem, we dissect it.
       Some foundation theory:
         1) 4 tuple TCP: TCP uses 4-tuple (source IP, source port, destination IP, destination port)
            1 tcp connect define form 4 parameter source IP + source port + destination IP + destination port,combined 4 parameter is unique. It's required for defined one connect
         2) ulimit in Linux: in Linux, everything is file, ulimit is parameter define maximum number of files open in one time
         3) design LB : endpoint ==> LB  ==> backend 
           
   2) Solution and ideal from 3 theory :
         1) From theory 1, don't use one IP address for load test. You have a problem limit port in one ip, if you up to up rqs and rqs later sometime. We choose cloud for test. In cloud, request from instance cloud to endpoint backend pass throw many DNS cloud, and list DNS change for every request, have change many ip always change. It's very good for load test.
         2) From theory 2, we set up ulimit to maximum in linux.
         3) From theory 3, in load test, we have data not simple is {"ok"} and response form backend very fast. But in rqs high, you pay a lot of money and pay many times for implements it. We have simple solution, we use big project for load test. EX: gooogle.com:80 ==> return 404, and response 60 character in 100 ms, it's perfect for load test. 
            But if we call google form one ip, google still block and status code change. We call google.com (backend) from aws (L7LB), it's very good and keep status response is 404.
   
   3) Result: <br/>
      1) in EC2_C5_4X_larger: >= 100 K rqs <br/>
          +) for LB: we build one instance EC2_C5_4X_larger for LB, we run docker with 1_LoadBalancing/L7Loadbalancer/LoadBalancerBenchmark/docker-compose-haproxy-local-to-gg.yml in this. <br/>
          +) for Load Test, we build  one instance EC2_C5_4X_larger, build docker and use bombardier for load test. It's same : docker run -ti --rm --ulimit nofile=65535:65535 --network=host alpine/bombardier --http1 -c 400   -d 600s -t 1s  -l http://ec2-13-250-40-69.ap-southeast-1.compute.amazonaws.com:8080/ <br/>
         =>>>>>>>>>>>>>>>>>>>  result:  <br/>
          +) we test with rqs from 1000 to 100.000 rqs, time test form 10 s to 10 min,  total request from 10 000 to 65080275, cpu <= 60 %, ram is very good  <br/>
          +) when we stop test, cpu is very fast to ~ 0 or 1 %, ram is very fast free => don't have leak ram, leak cpu  <br/>
          +) detail result in link 1_LoadBalancing/L7Loadbalancer/LoadBalancerBenchmark/result/c5_4x_large.md  <br/>
         =>>>>>>>>>>>>>>>>>>>>  Haproxy is very good for rqs, with one instance EC2_C5_4X_larger, we pass and run stable 100.000 rqs, we run stable in longtime, we have 100 M request continuous in 10 to 15 min and cpu and ram is good.  <br/>


todo:
config lam sao de request to one domain
not one service

point one domain with multiple ip addresses

auto update file config


./bench http://localhost:8080 ./benchmark-result/n1-highcpu-8-nginx 10s 300

wrk -t4 -c10 -d10s http://localhost:8000/

change nb thread, maxcount and test again
nbproc / nbthread ==> test again
https : test
*/  
*/