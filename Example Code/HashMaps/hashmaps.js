// This is the hashing function for the Hash Map
function hashStr(str) {
  let finalHash = 0;
  for (let i = 0; i < str.length; i++) {
    // returns the UTF-16 representation of the character at each index
    const charCode = str.charCodeAt(i);
    finalHash += charCode;
  }
  return finalHash;
}

console.log(hashStr("testkey"));

const m = new Hashmap();
m.set("name", "Jake");
console.log(m, get("name"));
