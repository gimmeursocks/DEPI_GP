FROM node:20-alpine

# Set working directory
WORKDIR /app

# Copy only package files first
COPY package*.json ./

# Install production dependecies only
RUN npm ci --omit=dev

# Copy rest of application files
COPY . .

# Best practice to use a non-root user
RUN addgroup -S app && adduser -S app -G app
USER app

# Expose the port
EXPOSE 8088

# Run the application
CMD ["npm", "start"]