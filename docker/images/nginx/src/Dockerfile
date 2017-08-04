FROM nginx:1.10

ENV TZ Asia/Shanghai
ENV nginx_home /etc/nginx

# remove default configure files
RUN rm -rf ${nginx_home}/conf.d
RUN rm -f ${nginx_home}/nginx.conf

# # Make my configure file
COPY nginx.conf ${nginx_home}/nginx.conf
#COPY conf.d/* ${nginx_home}/conf.d/
