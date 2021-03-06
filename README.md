# The Ray Tracer Challenge

Brice's attempt at the ray tracer challenge in Go

## Chapters

## Chapter 12 - Cubes

![Scene with a cube](output/chapter12.png)

### Chapter 11 – Reflection and Refraction

![Scene with a refractive surface](output/chapter11-refraction.png)
![Scene with a reflective surface](output/chapter11-reflection.png)


### Chapter 10 – Shaders

![Scene with shaders applied](output/chapter10.png)

### Chapter 9 – Planes

![Lit sphere on a plane](output/chapter9.png)

### Chapter 8 – Shadows

![Rendered scene with shadows](output/chapter8.png)
![Rendered scene with multiple lights](output/chapter8/chapter8-multilight.png)
![Animated render of shadows on a sphere](output/chapter8/animation/out.gif)

### Chapter 7 – Scenes

![Rendered scene](output/chapter7.png)

### Chapter 6 – Lighting

![Rendered sphere with lighting](output/chapter6.png)

### Chapter 5 – Object Transforms

![A transformed sphere](output/chapter5.png)

### Chapter 4 – Transform Matrices

![Transformation matrix example](output/chapter4.png)

## Notes

You can use [the Open Asset Importer (assimp)](https://github.com/assimp/assimp) to convert `.stl` files to `.obj` files.

It's available on Mac via [homebrew](https://brew.sh/). (`brew install assimp`)

```bash
$ assimp export Model.stl Model.obj
```

## TODO

- [ ] Implement cones
- [ ] Add some XYZ helper arrows
- [ ] Bounding boxes for efficiency (page 200)
- [ ] Named entities and scene search
- [ ] YAML loader for materials
- [ ] YAML external scene description
- [ ] Profile and optimise rendering function
- [ ] Orbit movement function
- [ ] UV Mapping for textures
- [ ] Optimise shaders with raw values types
- [ ] Transparency shadows
- [x] Parallelise rendering across mutliple threads
- [x] Progress indicator on render
- [x] Refactor Lights to be entities (so that they have a transform, potision, etc...)