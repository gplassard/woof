const { GoProject } = require('@gplassard/projen-extensions');

const project = new GoProject({
    name: 'woof',
});
project.addGitIgnore('dist');
project.synth();