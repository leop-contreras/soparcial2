FROM node:16 AS build

WORKDIR /front

COPY ./app/package.json ./app/package-lock.json ./
RUN npm install

COPY ./app ./
RUN npm run build

FROM nginx:alpine

#Corre mas rapido y se supone que es mejor. Si entendi bien, es react copilado?
COPY --from=build /front/build /usr/share/nginx/html
#RUN npm start

EXPOSE 80
EXPOSE 3000
