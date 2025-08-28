#!/bin/bash

# Script pour tester l'API de conversion

API_URL="http://localhost:8080"

echo "🧪 Testing DevOps Converter API..."

# Test 1: Health check
echo -e "\n1️⃣ Testing health endpoint..."
curl -s "$API_URL/health/" | jq . 2>/dev/null || echo "Health endpoint test failed"

# Test 2: Get available converters
echo -e "\n2️⃣ Testing converters endpoint..."
curl -s "$API_URL/api/v1/info/converters" | jq . 2>/dev/null || echo "Converters endpoint test failed"

# Test 3: Validate docker-compose file
echo -e "\n3️⃣ Testing validation endpoint..."
cat << 'EOF' > /tmp/test-validation.json
{
  "type": "docker-compose",
  "content": "version: '3.8'\nservices:\n  web:\n    image: nginx:latest\n    ports:\n      - \"80:80\""
}
EOF

curl -s -X POST "$API_URL/api/v1/convert/validate" \
  -H "Content-Type: application/json" \
  -d @/tmp/test-validation.json | jq . 2>/dev/null || echo "Validation test failed"

# Test 4: Convert docker-compose to Kubernetes
echo -e "\n4️⃣ Testing conversion endpoint..."
cat << 'EOF' > /tmp/test-conversion.json
{
  "type": "docker-compose",
  "content": "version: '3.8'\nservices:\n  web:\n    image: nginx:latest\n    ports:\n      - \"80:80\"\n    environment:\n      - ENV=production",
  "options": {
    "namespace": "test",
    "serviceType": "LoadBalancer"
  }
}
EOF

curl -s -X POST "$API_URL/api/v1/convert/" \
  -H "Content-Type: application/json" \
  -d @/tmp/test-conversion.json | jq . 2>/dev/null || echo "Conversion test failed"

# Cleanup
rm -f /tmp/test-validation.json /tmp/test-conversion.json

echo -e "\n✅ API tests completed!"
