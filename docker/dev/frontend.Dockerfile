FROM node:22-alpine
WORKDIR /app

COPY trusmi_fe/package*.json ./
RUN npm install

COPY trusmi_fe/ ./

EXPOSE 5173
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0", "--port", "5173"]
