#! /usr/bin/env node
import { replaceInFileSync } from "replace-in-file";

const oldVersion = process.argv[2];
const newVersion = process.argv[3];

console.log(`Bumping version from ${oldVersion} to ${newVersion}`);

replaceInFileSync({
    files: "./**/package.json",
    from: new RegExp(`"${oldVersion}"`, "g"),
    to: `"${newVersion}"`,
});

replaceInFileSync({
    files: "./**/package.json",
    from: new RegExp(`"\\^${oldVersion}"`, "g"),
    to: `"^${newVersion}"`,
});
