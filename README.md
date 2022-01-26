# Cut-Films-Into-Scenes-And-Stream-The-Final-Film

You are working on a project and need to cut films into scenes. to help streamline the
creation of the final films, the team needs to develop an automated way of breaking up individual shots (short sequence  from a particular camera angle) in a film into scenes (a sequence of shots). there is already an algorithm that breaks the film up into shots and labels them with a a letter. Identical shots are labelled with the same letter. write a function to split the film into as many short scenes as possible without confusing viewers by having the same shot appear in different scenes. This will partition the sequence of shot labels into scenes so that no shot label appears in more than one scene and each scene is as short as possible. the output should be the length of each scene.

Input
the input to the function/method consists of an argument - inputList, a list of characters 
representing the sequence of shots.

Output 
Return a list of integers representing the length of each scene, in the order in which it
appears in the given sequence of shots.

Examples:
example 1:
input :
inputList = [a,b,a,b,c,b,a,c,a,d,e,f,e,g,d,e,h,i,j,h,k,l,i,j]
output:
[9,7,8]

Explanation:
the first scene consists of the shots a,b and c.
the second scene consists of d,e,f and g.
finally, the last scene consists of h,i,j,k and l.
the answer is 9,7,8 because a,b, and c only appear in the 
first 9 characters, then d,e,f, and g appear in the next 7.
the final 8 characters consist entirely of h,i,j,k, and l.

example 2:
input:
inputList=[a,b,c]
output:
[1,1,1]

explanation:
because there are no recurring shots, all shots can be in
the minimal length 1 scene.

example 3
input:
inputList =[a,b,c,a]
output:
[1]
explanation:
because 'a' appears more than once, everything between the first and last appearance of 'a' must be in the same list.

Testcase 1:
[a,b,c,d,a,e,f,g,h,i,j,e]

expected return values:
5 7

Testcase 2:
input:
[z,w,c,b,z,c,h,f,i,h,i]

expected return values:
6 5




