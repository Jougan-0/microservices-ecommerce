FROM node:18

WORKDIR /app

ARG NEXT_PUBLIC_API_URL
ENV NEXT_PUBLIC_API_URL=$NEXT_PUBLIC_API_URL

COPY package.json package-lock.json ./
RUN npm install

COPY . .

RUN NEXT_PUBLIC_API_URL=$NEXT_PUBLIC_API_URL npm run build

CMD ["npm", "run", "start"]
