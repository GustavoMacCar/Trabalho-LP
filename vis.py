from matplotlib import pyplot as plt
import pandas as pd

ds = pd.read_csv("squares.csv")

plt.scatter(ds[ds['colour']=='orange']['x'],ds[ds['colour']=='orange']['y'],c='orange')
plt.scatter(ds[ds['colour']=='blue']['x'],ds[ds['colour']=='blue']['y'],c='b')
plt.show()