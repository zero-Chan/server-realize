FROM ruby:2.3.1

ENV TZ Asia/Shanghai

RUN gem sources --add https://ruby.taobao.org/ --remove https://rubygems.org/
RUN gem install redis

ADD ./redis-trib.rb /redis-trib.rb

ENTRYPOINT ["/redis-trib.rb"]
