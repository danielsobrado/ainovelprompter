FROM node:14

WORKDIR /app

COPY package*.json ./

RUN npm install
RUN npm install react-router-dom axios @mui/material @emotion/react @emotion/styled @mui/icons-material react-toastify

COPY . .
RUN npm run build

EXPOSE 3000

CMD ["npm", "start"]