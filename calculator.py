from math import pi

"""This program finds the area of a circle when given it's radius"""

radius = int(input("What is the radius of your circle?"))
area = pi * radius * radius
area = round(area, 2)  # area to 2 decimal places

print("The area of your circle is ", area)