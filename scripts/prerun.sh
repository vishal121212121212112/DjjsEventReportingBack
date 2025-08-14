#!bin/bash


create_swagger(){
   echo "creating swagger file"
  # for creating swagger file
   swagger-cli bundle docs/main.yaml -o docs/swagger.yaml -t yaml

   echo "swagger file has been created."

}

create_swagger