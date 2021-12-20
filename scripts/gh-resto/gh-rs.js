#!/usr/bin/env node

const fs = require("fs");
const execa = require("execa");
const path = require("path");
const rm = require("rimraf");
const mkdirp = require("mkdirp");
const sh = require("shelljs");

const VERSION_CMD = sh.exec("git describe --abbrev=0 --tags")
const VERSION_DATE_CMD = sh.exec("cat date.txt")

const VERSION = VERSION_CMD.replace("\n", "").replace("\r", "");
const VERSION_DATE = VERSION_DATE_CMD.replace("\n", "").replace("\r", "");

const ROOT = __dirname;
const TEMPLATES = path.join(ROOT, "templates");

async function updateRestoExtension(ghRestoDir) {
  const templatePath = path.join(TEMPLATES, "gh-resto");
  const template = fs.readFileSync(templatePath).toString("utf-8");

  const templateReplaced = template
    .replace("CLI_VERSION", VERSION)
    .replace("CLI_VERSION_DATE", VERSION_DATE)

  fs.writeFileSync(
    path.join(ghRestoDir, "gh-resto"),
    templateReplaced
  );
}

async function updateExtension() {
  const tmp = path.join(__dirname, "tmp");
  const extensionDir = path.join(tmp, "gh-resto");
  mkdirp.sync(tmp);
  rm.sync(extensionDir);

  console.log(
    `cloning https://github.com/abdfnx/gh-resto to ${extensionDir}`
  );

  await execa("git", [
    "clone",
    "https://github.com/abdfnx/gh-resto.git",
    extensionDir,
  ]);

  console.log(`done cloning abdfnx/gh-resto to ${extensionDir}`);

  console.log("updating local git...");

  await updateRestoExtension(extensionDir);

  const git = async (args, opts = {}) => {
    await execa("git", ["-C", extensionDir, ...args], opts);
  };

  await git(["add", "gh-resto"]);
  await git(["config", "--local", "core.pager", "cat"]);
  await git(["diff", "--cached"], { stdio: "inherit" });
  await git(["commit", "-m", "resto `" + VERSION + "`"]);

  if (process.env.SKIP_GIT_PUSH === undefined) {
    await git(["push", "origin", "main"]);
  }
}

updateExtension().catch((err) => {
  console.error(`error running scripts/gh-resto/gh-rs.js`, err);
  process.exit(1);
});
