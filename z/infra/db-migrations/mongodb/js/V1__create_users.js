db = db.getSiblingDB('sandbox_db');

db.createCollection('users', {
  validator: {
    $jsonSchema: {
      bsonType: 'object',
      required: ['name', 'email', 'password', 'created_at', 'updated_at'],
      properties: {
        name: {
          bsonType: 'string',
          maxLength: 100,
          description: 'must be a string and is required'
        },
        email: {
          bsonType: 'string',
          maxLength: 150,
          description: 'must be a string and is required'
        },
        password: {
          bsonType: 'string',
          maxLength: 255,
          description: 'must be a string and is required'
        },
        created_at: {
          bsonType: 'date',
          description: 'must be a date and is required'
        },
        updated_at: {
          bsonType: 'date',
          description: 'must be a date and is required'
        }
      }
    }
  }
});

db.users.createIndex({ email: 1 }, { unique: true });
db.users.createIndex({ created_at: -1 });