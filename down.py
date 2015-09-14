def download(plan=True,zastepstwa=True):
    import urllib2
    we = urllib2.urlopen('http://loiv.torun.pl/index.php/pl/dla-uczniow/organizacja-zajec/plan-lekcji').read()
    import re
    pse = re.search(r'attachments.*plan_KLAS.pdf',we).group().split('"')
    zse = re.search(r'attachments.*ast.*pstwa.*pdf',we).group().split('"')
    listaza = set()
    listape = set()
    zl = set()
    for i in zse:
        try: yu = re.search(r'attachments.*ast.*pstwa.*pdf',i,re.S).group()
        except AttributeError: continue
        listaza.add(yu)
    for i in pse:
        if i is not None:
            vbnm = re.search(r'attachments.*plan_KLAS.pdf',i,re.S)
            if vbnm is None: continue
            listape.add(vbnm.group())
    print pse,zse,listaza
    if plan:
        oj = listape.pop()
        try: p = urllib2.urlopen('http://loiv.torun.pl/'+urllib2.quote(oj)).read()
        except urllib2.HTTPError: print oj, listape ; raise
    if zastepstwa:
        for i in listaza:
            try: zl.add(urllib2.urlopen('http://loiv.torun.pl/'+urllib2.quote(i)).read())
            except urllib2.HTTPError: print i, listaza ; raise
    return {'p':p,'z':zl}

