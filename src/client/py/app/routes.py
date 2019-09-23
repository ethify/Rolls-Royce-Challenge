from flask import Flask, render_template, session, redirect, url_for, request
from app import app
from os import listdir, system, chdir
from os.path import isfile, join, abspath, dirname
import subprocess
#from hfc.fabric import Client

jspath = abspath(join(dirname( __file__ ), '..', '..','javascript'))

@app.context_processor
def members():
	global wallet_path
	global onlyfiles
	wallet_path = abspath(join(jspath,'wallet'))
	onlyfiles = [f for f in listdir(wallet_path) if not isfile(join(wallet_path, f))]
	return dict(members=onlyfiles)

@app.route('/index')
def index():
	return "Index"

@app.route('/wri-list')
def wri_list():
	if 'wallet' in session:
		user=session['wallet']
		wri_list = []
		chdir(jspath)
		count = subprocess.check_output(["node", "queryReqCount.js", session['wallet']])
		count = int(count.decode())
		for c in range(1, count+1):
			a = subprocess.check_output(["node", "queryGetReq.js", session['wallet'], "req"+str(count)])
			wri_list.append(a.decode())
			print(wri_list)
		return render_template('wri.html',all_polls=wri_list, prev_polls={})
	return redirect(url_for('signin'))

@app.route('/wi-list')
def wi_list():
	if 'wallet' in session:
		user=session['wallet']
		wri_list = []
		chdir(jspath)
		count = subprocess.check_output(["node", "queryWorkCount.js", session['wallet']])
		count = int(count.decode())
		for c in range(1, count+1):
			a = subprocess.check_output(["node", "queryGetWork.js", session['wallet'], "work"+str(count)])
			wri_list.append(a.decode())
			print(wri_list)
		return render_template('wi.html',all_polls=wri_list, prev_polls={})
	return redirect(url_for('signin'))

@app.route('/wri-create')
def forms():
	if 'wallet' in session:
		return render_template('forms.html')
	return redirect(url_for('signin'))

@app.route('/wi-create')
def forms2():
	if 'wallet' in session:
		return render_template('forms2.html')
	return redirect(url_for('signin'))

@app.route('/admin', methods=['GET','POST'])
def admin():
	if request.method == 'POST':
		name = request.form['name']
		chdir(jspath)
		system("node registerUser.js " + name)
	if 'wallet' in session and session['wallet'] == 'admin':
		return render_template('admin.html')
	return redirect(url_for('signin'))

@app.route('/signup')
def signup():
	return render_template('signup.html')

@app.route('/signin')
def signin():
	return render_template('login.html')
