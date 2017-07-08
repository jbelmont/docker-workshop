'use strict'

const express = require('express')
const path = require('path')
const fs = require('fs')
const multipart = require('connect-multiparty')
const fsExtra = require('fs-extra')

const app = express()
const multipartMiddleware = multipart()

app.use(express.static(path.join(__dirname, 'public')))
app.get('/', (req, res) => {
  res.sendFile(path.join(__dirname, 'index.html'))
})

app.post('/', multipartMiddleware, (req, res) => {
  if (req.files && req.files.file && req.files.file.originalFilename) {
    fsExtra.copySync(req.files.file.path, path.join(__dirname, 'public', req.files.file.originalFilename))
  }
  res.sendFile(path.join(__dirname, 'index.html'))
})

app.get('/files', (req, res, cb) => {
  fs.readdir(path.join(__dirname, 'public'), (err, items) => {
    if (err) {
      return cb(err)
    }
    res.statusCode = 200
    res.json(items)
  })
})

app.listen(3004, () => {
  console.log('Server listening on http://localhost:3004/')
})
