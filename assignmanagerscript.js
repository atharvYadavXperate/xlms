const BASE_URL = "http://localhost:8080"; // change if needed
const LIMIT = 50;
const ROOT_ADMIN_ID = 1;

// % of USERS that will NOT get any manager
const UNASSIGNED_USER_PERCENTAGE = 20;

function shouldSkipUser() {
  return Math.random() * 100 < UNASSIGNED_USER_PERCENTAGE;
}

async function fetchJSON(url, options = {}) {
  const res = await fetch(url, {
    method: options.method || "GET",
  });

  if (!res.ok) {
    const text = await res.text();
    throw new Error(`${res.status} ${text}`);
  }

  return res.json();
}

async function assignManagersHierarchy() {
  let page = 1;

  const admins = [];
  const managers = [];
  const users = [];

  // 🔹 STEP 1: FETCH ALL USERS
  while (true) {
    const data = await fetchJSON(
      `${BASE_URL}/users/getusers?page=${page}&limit=${LIMIT}`
    );

    const list = data.data;

    if (!list || list.length === 0) break;

    for (const user of list) {
      if (user.id === ROOT_ADMIN_ID) continue;

      if (user.role_id === 1) admins.push(user);
      else if (user.role_id === 2) managers.push(user);
      else if (user.role_id === 3) users.push(user);
    }

    if (list.length < LIMIT) break;
    page++;
  }

  console.log(
    `Fetched → Admins: ${admins.length}, Managers: ${managers.length}, Users: ${users.length}`
  );

  // 🔹 STEP 2: ROOT ADMIN → ADMINS
  for (const admin of admins) {
    try {
      await fetchJSON(
        `${BASE_URL}/managers/set?manager=${ROOT_ADMIN_ID}&user=${admin.id}`,
        { method: "POST" }
      );

      console.log(`✅ ROOT_ADMIN → ADMIN ${admin.id}`);
    } catch (err) {
      console.error(`❌ ADMIN ${admin.id}`, err.message);
    }
  }

  // 🔹 STEP 3: ADMINS → MANAGERS (round-robin)
  let adminIndex = 0;

  for (const manager of managers) {
    const admin = admins[adminIndex % admins.length];
    adminIndex++;

    if (!admin) {
      console.warn("⚠️ No admins available for managers");
      break;
    }

    try {
      await fetchJSON(
        `${BASE_URL}/managers/set?manager=${admin.id}&user=${manager.id}`,
        { method: "POST" }
      );

      console.log(`✅ ADMIN ${admin.id} → MANAGER ${manager.id}`);
    } catch (err) {
      console.error(`❌ MANAGER ${manager.id}`, err.message);
    }
  }

  // 🔹 STEP 4: MANAGERS → USERS
  let managerIndex = 0;

  for (const user of users) {
    if (shouldSkipUser()) {
      console.log(`⏭️ USER ${user.id} left unassigned`);
      continue;
    }

    const manager = managers[managerIndex % managers.length];
    managerIndex++;

    if (!manager) {
      console.warn("⚠️ No managers available for users");
      break;
    }

    try {
      await fetchJSON(
        `${BASE_URL}/managers/set?manager=${manager.id}&user=${user.id}`,
        { method: "POST" }
      );

      console.log(`✅ MANAGER ${manager.id} → USER ${user.id}`);
    } catch (err) {
      console.error(`❌ USER ${user.id}`, err.message);
    }
  }

  console.log("🎉 Manager hierarchy assignment completed");
}

// RUN
assignManagersHierarchy().catch(err =>
  console.error("❌ Script failed:", err.message)
);
