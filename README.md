encore-cmd is meant to explore a script-based configuration management tool.

It uses a cmdfile akin to Dockerfile syntax. However, it only provides a limited functionality since we want to separate configuration management from system administration tasks.

A cmdfile looks like the following:

GO hostenv PATH
ENV MYVARIABLE myvar
GO getenv MYVARIABLE

RUN apt-key adv --keyserver pgp.mit.edu --recv-keys 573BFD6B3D8FBC641079A6ABABF5BD827BD9BF62
RUN touch /etc/apt/sources.list
RUN echo "deb http://nginx.org/packages/mainline/debian/ wheezy nginx" > /etc/apt/sources.list
RUN apt-get update 
RUN apt-get install -y nginx
RUN rm -rf /var/lib/apt/lists/*

To run,

encore-cmd /path/to/cmdfile
