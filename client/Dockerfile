FROM node:18

WORKDIR /app

COPY package*.json ./

RUN npm install

ENV NODE_OPTIONS=--openssl-legacy-provider

COPY . .
RUN npm run build

EXPOSE 3000

CMD ["npm", "start"]