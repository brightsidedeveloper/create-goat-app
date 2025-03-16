#!/usr/bin/env node

const fs = require('fs-extra')
const path = require('path')
const { program } = require('commander')
const inquirer = require('inquirer')
const { execSync } = require('child_process') // For running shell commands

program
  .version('1.0.0')
  .argument('[project-name]', 'Name of the project directory')
  .description('Create a new Goat App')
  .action(async (projectName) => {
    // Prompt for project name if not provided
    const answers = await inquirer.prompt([
      {
        type: 'input',
        name: 'projectName',
        message: 'Enter project name:',
        default: 'goat-app',
        when: !projectName, // Only ask if projectName isn’t provided
      },
    ])

    // Use provided projectName or the prompted one
    const finalProjectName = projectName || answers.projectName
    const targetDir = path.join(process.cwd(), finalProjectName)

    // Check if directory exists
    if (fs.existsSync(targetDir)) {
      console.error(`Error: Directory ${finalProjectName} already exists!`)
      process.exit(1)
    }

    // Create the directory
    fs.mkdirSync(targetDir)

    // Copy template
    const templateDir = path.join(__dirname, 'template')
    try {
      fs.copySync(templateDir, targetDir)
    } catch (err) {
      console.error('Error copying template:', err)
      process.exit(1)
    }

    console.log(`Setting up ${finalProjectName}...`)

    // Automatically run npm install in the target directory
    try {
      console.log('Installing dependencies with npm install...')
      execSync('npm install', { cwd: targetDir, stdio: 'inherit' })
    } catch (err) {
      console.error('Error running npm install:', err)
      process.exit(1)
    }

    // Change into the wasm subdirectory and run go mod tidy
    const wasmDir = path.join(targetDir, 'wasm')
    if (fs.existsSync(wasmDir)) {
      try {
        console.log('Running go mod tidy in wasm directory...')
        execSync('go mod tidy', { cwd: wasmDir, stdio: 'inherit' })
      } catch (err) {
        console.error('Error running go mod tidy:', err)
        process.exit(1)
      }
    } else {
      console.warn('Warning: wasm directory not found in template. Skipping go mod tidy.')
    }

    console.log(`Success! Created and set up ${finalProjectName} at ${targetDir}`)
    console.log('To get started:')
    console.log(`  cd ${finalProjectName}`)
    console.log('  make dev')
  })

program.parse(process.argv)
