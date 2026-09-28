import http from 'node:http';
import fs from 'node:fs';
import path from 'node:path';
import { fileURLToPath } from 'node:url';
const root=path.dirname(fileURLToPath(import.meta.url));
const mime={'.html':'text/html; charset=utf-8','.css':'text/css; charset=utf-8','.js':'text/javascript; charset=utf-8'};
http.createServer((req,res)=>{
 const url=new URL(req.url,'http://localhost');
 let file;
 if(url.pathname==='/slides') file=path.join(root,'presentation/index.html');
 else if(url.pathname==='/app'||url.pathname==='/') file=path.join(root,'app/index.html');
 else if(url.pathname.startsWith('/presentation-assets/')) file=path.join(root,'presentation/assets',path.basename(url.pathname));
 else if(url.pathname.startsWith('/app-assets/')) file=path.join(root,'app/assets',path.basename(url.pathname));
 else file=path.join(root,'index.html');
 fs.readFile(file,(err,data)=>{if(err){res.writeHead(404);res.end('Not found');return;}res.writeHead(200,{'Content-Type':mime[path.extname(file)]||'application/octet-stream'});res.end(data);});
}).listen(8082,()=>console.log('UOB Library preview at http://localhost:8082/app'));
