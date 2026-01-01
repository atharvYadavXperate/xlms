const API_URL = "http://localhost:8080/users/register";

// config
const TOTAL_USERS = 200;
const ADMINS = 5;
const MANAGERS = 15;

// helper to create random email
function randomEmail(prefix, i) {
  return `${prefix}${i}_${Date.now()}@mail.com`;
}

// build users array
const users = [];

// admins
for (let i = 1; i <= ADMINS; i++) {
  users.push({
    full_name: `Admin ${i}`,
    email: randomEmail("admin", i),
    role_id: 1,
  });
}

// managers
for (let i = 1; i <= MANAGERS; i++) {
  users.push({
    full_name: `Manager ${i}`,
    email: randomEmail("manager", i),
    role_id: 2,
  });
}

// normal users
for (let i = 1; i <= TOTAL_USERS - ADMINS - MANAGERS; i++) {
  users.push({
    full_name: `User ${i}`,
    email: randomEmail("user", i),
    role_id: 3,
  });
}

async function createUsers() {
  let success = 0;

  for (const user of users) {
    try {
      const res = await fetch(API_URL, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(user),
      });

      if (!res.ok) {
        const text = await res.text();
        console.error("❌ Failed:", user.email, text);
        continue;
      }

      success++;
      console.log(`✅ Created (${success}/${users.length}):`, user.email);
    } catch (err) {
      console.error("❌ Error:", user.email, err.message);
    }
  }

  console.log(`\n🎉 Done! ${success} users created.`);
}

createUsers();
