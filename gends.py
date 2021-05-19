import numpy as np
import pandas as pd
n=1000
x = np.random.rand(n)*20-10
y = np.random.rand(n)*20-10
col = np.repeat(["orange"],n)
col[(x+3)*(y-5)>0]="blue"
ds = pd.DataFrame(np.array([x,y,col]).T,columns=["x","y","colour"])
ds.to_csv("x3y5.csv")