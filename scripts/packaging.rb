#!/usr/bin/env ruby

require 'erb'

# Target platforms and architectures
TARGETS = [
  'linux'   => {
    'deb' => ['i386', 'amd64'],
    'rpm' => ['i386', 'amd64'],
  },
  'darwin'  => {
    'pkg' => ['i386', 'amd64']
  },
]

# Rendering PackageInfo
class PackageInfo
  def initialize(version, content_path)
    @version = version
    @template = File.read("scripts/PackageInfo.erb")
    @size = directory_size(content_path)
  end

  def render
    ERB.new(@template).result(binding)
  end

  def save(file)
    File.open(file, "w+") do |f|
      f.write(render)
    end
  end

  # Return directory size in KBytes
  def directory_size(path)
    size=0
    Dir.glob(File.join(path, '**', '*')) { |file| size+=File.size(file) }
    return (size/1024)
  end
end

# Build targeted binaries
def build(package, type, version, os, arch)
  # Dependencies
  go_arch  = arch

  if arch == 'i386'
    go_arch = '386'
  end

  go_build_target = "#{package[:path]}/#{package[:name]}-#{os}-#{arch}"
  sh %{CGO_ENABLED=0 GOARCH=#{go_arch} GOOS=#{os} go build -o #{go_build_target}}
  File.chmod(0755, go_build_target)

  if os == "darwin"
    root_dir = "root-pkg"
    pkg_path = "#{package[:path]}/tekleader.pkg"
    bin_path = "#{package[:path]}/#{root_dir}/usr/local/bin"

    mkdir_p bin_path
    mkdir_p pkg_path

    # Generate the payload
    cp go_build_target, "#{bin_path}/tekleader"
    system("( cd #{package[:path]}/#{root_dir} && find . | cpio -o --format odc --owner 0:80 | gzip -c ) > #{pkg_path}/Payload")

    # Generate the package description
    pkg_info = PackageInfo.new(version, "#{package[:path]}/#{root_dir}")
    pkg_info.save("#{pkg_path}/PackageInfo")

    # Generate the Bill Of Materials
    system("mkbom -u 0 -g 80 #{package[:path]}/#{root_dir} #{pkg_path}/Bom")

    # Build the resulting pkg
    system("( cd #{pkg_path} && xar --compression none -cf \"../#{package[:name]}-#{version}-#{arch}.pkg\" * )")

    # Clean
    rm_r(pkg_path, :force => true)
    rm_r("#{package[:path]}/#{root_dir}", :force => true)
  else
    sh %W{fpm
      --force
      -s dir
      -t #{type}
      -a #{arch}
      -n #{package[:name]}
      -p #{package[:path]}
      --maintainer "#{package[:maintainer]}"
      --description "#{package[:info]}"
      --license "#{package[:licence]}"
      --url "#{package[:url]}"
      --vendor "#{package[:vendor]}"
      --version "#{version}"
      --deb-use-file-permissions
      --rpm-use-file-permissions
      #{go_build_target}=/usr/local/bin/#{package[:name]}
    }.join(' ')
  end
end

def release(package, version)
  TARGETS.each do |target|
    target.each do |os, pkgmap|
      pkgmap.each do |pkg, archs|
        archs.each do |arch|
          build(package, pkg, version, os, arch)
        end
      end
    end
  end
end
