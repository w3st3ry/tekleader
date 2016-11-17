#!/usr/bin/env ruby

require 'git'
require 'logger'
require 'rake/testtask'
require 'rubygems'
require_relative 'scripts/packaging'

package = {
  :name       => 'tekleader',
  :maintainer => 'valentin.pichard@corp.ovh.com',
  :info       => 'Teakleader makes it possible some cool things over EPITECH intranet api',
  :url        => 'https://intra.epitech.eu/',
  :licence    => 'MIT',
  :path       => 'releases',
}

# -----------------------------------------
#  PREPARE-RELEASE
# -----------------------------------------
desc 'Prepare the releasing processs'
task :prepare_release do |t, args|
  system("gem install bundler")
  system("bundle install")
  system("go get github.com/aktau/github-release")
end

# -----------------------------------------
#  RELEASE
# -----------------------------------------
desc 'Release a new version'
task :release, [:version] => [:prepare_release] do |t, args|
  path = package[:path]
  tag = "v#{args.version}"

  # Push on master
  g = Git.open("#{ENV['GOPATH']}/src/github.com/w3st3ry/tekleader", :log => Logger.new(STDOUT))
  g.checkout(:master)
  g.add(['RELEASE.md', 'tekleader/version.go'])
  g.commit(["Release #{args.version}",
    "",
    "Signed-off-by: #{g.config('user.name')} <#{g.config('user.email')}>",
    "",
  ].join("\n"))
  g.push(:origin, :master)

  # Merge to release branch and push
  g.checkout(:release)
  g.merge(:master)
  g.add_tag(tag)
  g.push(:origin, :release)
  g.push(:origin, :release, {:tags => true})

  # Release
  mkdir_p path
  release(package, args.version)

  description = File.read("RELEASE.md")
  system("github-release release --user w3st3ry --repo tekleader --tag #{tag} --name \"Version #{args.version}\" --description \"#{description}\" --pre-release")

  Dir["#{path}/*"].each do |file|
    system("github-release upload --user w3st3ry --repo tekleader --tag #{tag} --name #{File.basename(file)} --file #{file}")
    File.delete(file)
  end

  g.checkout(:master)
end
